package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"mt2is/pkg/category"
	"net/http"
	"sort"

	"github.com/gorilla/sessions"
)

//CategoryHandler is a HTTP Handler which displays items in categories
type CategoryHandler struct {
	catTree  *category.NodeTree
	template *template.Template
}

//CompileTemplates parses the main template and all associated templates. They are stored in memory afterwards
func (c *CategoryHandler) CompileTemplates() error {
	tpl, err := template.ParseFiles("tpl/base.go.html", "tpl/pages/category.go.html")
	if err != nil {
		return fmt.Errorf("Failed compiling templates: %v", err)
	}
	c.template = tpl
	return nil
}

//NewCategoryHandler instantiates and returns a ready to use handler for category page
func NewCategoryHandler(catProvider category.NodeTreeProvider) (*CategoryHandler, error) {
	catTree, err := catProvider.Provide()
	if err != nil {
		return nil, fmt.Errorf("Can't instantiate the category handler: %v", err)
	}
	h := &CategoryHandler{catTree: catTree}
	err = h.CompileTemplates()
	if err != nil {
		return nil, fmt.Errorf("Can't instantiate the category handler: %v", err)
	}

	return h, nil
}

//ServeHTTP receives the requets and writes the executed template using the given writer
func (c *CategoryHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	roots := c.catTree.Roots()

	store := sessions.NewFilesystemStore("var/sessions", []byte("secret key"))

	sess, _ := store.Get(request, "prova2")
	fmt.Println(sess.Values["prova"])
	sess.Values["prova"] = "ciao2"
	sess.Save(request, writer)
	c.template.Execute(writer, roots)
	fmt.Fprintln(writer, request.Context().Value(accIDKey))
}

//SQLNodeTreeProvider provides a multi-root hierarchical tree made of categories
type SQLNodeTreeProvider struct {
	db *sql.DB
}

/*Provide is a concrete implementation of the Provide() method which returns a multi-root hierarchical tree.
 *The building process is made of two steps.
 *
 *The first one involves iterating over the SQL results,
 *building each node and saving linkage information for the next step.
 *The linkage information consists in each SQL row having a foreign key 'parent_id' which references its parent
 *
 *In the second step each node built previously is iterated and the linkage between other nodes is established.
 *At the same time nodes which are roots (has no parent) are filtered for the latter return
 */
func (p *SQLNodeTreeProvider) Provide() (*category.NodeTree, error) {
	rows, err := p.db.Query("SELECT category_id, name, trailer, description, link_segment, parent_id, priority FROM itemshop.categories")

	if err != nil {
		return nil, fmt.Errorf("Can't fetch the categories: %v", err)
	}

	//Linkage information
	idToNode := make(map[int]*category.Node)
	idToParentID := make(map[int]int)
	idToChildIds := make(map[int][]int)

	for rows.Next() {
		var parentID int
		var priority int
		cat := &category.Category{}
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.Trailer, &cat.Description, &cat.LinkSegment, &parentID, &priority); err != nil {
			return nil, fmt.Errorf("Can't map sql row to category: %v", err)
		}

		node := category.NewNode(cat, priority)

		idToNode[cat.ID] = node
		if parentID > 0 {
			idToParentID[cat.ID] = parentID
			idToChildIds[parentID] = append(idToChildIds[parentID], cat.ID)
		}
	}

	roots := make([]*category.Node, 0)
	for id, node := range idToNode {
		var parentNode *category.Node
		if parentID, ok := idToParentID[id]; ok {
			parentNode = idToNode[parentID]
			node.SetParent(parentNode)
		}

		if childIds, ok := idToChildIds[id]; ok {
			for _, childID := range childIds {
				if parentNode != nil && parentNode == idToNode[childID] {
					return nil, fmt.Errorf("Circular reference, node id %d has as parent and as a child the node id %d", id, childID)
				}
				node.AddChild(idToNode[childID])
			}
			node.SortChildren()
			if node.IsRoot() {
				roots = append(roots, node)
			}
		}
	}

	sort.Slice(roots, func(i, j int) bool { return roots[i].Priority < roots[j].Priority })

	catTree := category.NewNodeTree(roots)
	return catTree, nil
}
