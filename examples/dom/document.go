// +build js,wasm

package main

import (
	"fmt"
	"github.com/sternix/wasm"
)

func TestDocument(doc wasm.Document) {
	fmt.Printf("Document.URL: %s\n", doc.URL())
	fmt.Printf("Document.DocumentURI: %s\n", doc.DocumentURI())
	fmt.Printf("Document.Origin: %s\n", doc.Origin())
	fmt.Printf("Document.CompatMode: %s\n", doc.CompatMode())
	fmt.Printf("Document.CharacterSet: %s\n", doc.CharacterSet())
	fmt.Printf("Document.ContentType: %s\n", doc.ContentType())

	dt := doc.DocType()
	if dt == nil {
		errx("Document.DocumentType == NULL")
	}

	TestDocType(dt)

	TestLoginForm(doc)
}

func TestDocType(dt wasm.DocumentType) {
	TestNode(dt)
	fmt.Printf("DocumentType.Name: %s\n", dt.Name())
	fmt.Printf("Document.PublicId: %s\n", dt.PublicId())
	fmt.Printf("Document.SystemId: %s\n", dt.SystemId())
}

func TestNode(node wasm.Node) {
	fmt.Println("-----------------------------------")
	fmt.Println("Node Tests Started")

	fmt.Printf("Node.Type: %v\n", node.NodeType())
	fmt.Printf("Node.NodeName: %s\n", node.NodeName())
	fmt.Printf("Node.BaseURI: %s\n", node.BaseURI())
	fmt.Printf("Node.IsConnected: %t\n", node.IsConnected())
	if node.OwnerDocument() == nil {
		fmt.Println("Node.OwnerDocument == NULL")
	}

	if node.RootNode() == nil {
		fmt.Println("Node.RootNode == NULL")
	}

	if node.RootNode(wasm.RootNodeOptions{Composed: false}) == nil {
		fmt.Println("Node.RootNode with options false  == NULL")
	}
	if node.RootNode(wasm.RootNodeOptions{Composed: true}) == nil {
		fmt.Println("Node.RootNode with options false  == NULL")
	}

	if node.ParentNode() == nil {
		fmt.Println("Node.ParentNode == NULL")
	}

	if node.ParentElement() == nil {
		fmt.Println("Node.ParentElement == NULL")
	}

	fmt.Printf("Node.HasChildNodes: %t\n", node.HasChildNodes())
	if node.FirstChild() == nil {
		fmt.Println("Node.FirstChild == NULL")
	}

	if node.LastChild() == nil {
		fmt.Println("Node.LastChild == NULL")
	}

	if node.PreviousSibling() == nil {
		fmt.Println("Node.PreviousSibling == NULL")
	}

	if node.NextSibling() == nil {
		fmt.Println("Node.NextSibling == NULL")
	}

	/*
		nv := "xyz"
		node.SetNodeValue("xyz")

		if node.NodeValue() != nv {
			errx("Node.NodeValue != Node.SetNodeValue")
		}
	*/
	fmt.Printf("Node.NodeValue: %s\n", node.NodeValue())

	/*
		ntc := "klmn"
		node.SetTextContent(ntc)
		if node.TextContent() != ntc {
			errx("Node.TextContent != Node.SetTextContent")
		}
	*/
	fmt.Printf("Node.TextContent: %s\n", node.TextContent())

	node.Normalize()

	cn := node.CloneNode()
	if cn == nil {
		errx("Node.CloneNode == NULL")
	}

	cn = node.CloneNode(true)
	if cn == nil {
		errx("Node.CloneNode true == NULL")
	}

	if !node.IsEqualNode(cn) {
		fmt.Println("Node.IsEqualNode == FALSE")
	}

	if !node.IsSameNode(cn) {
		fmt.Println("Node.IsSameNode == FALSE")
	}

	PrintDocumentPosition(node.CompareDocumentPosition(cn))

	fmt.Printf("Node.Contains: %t\n", node.Contains(cn))

	// TODO
	//node.LookupPrefix()
	//node.LookupNamespaceURI()
	//node.IsDefaultNamespace()

	fmt.Println("Node Test Finished")
	fmt.Println("-----------------------------------")
}

func PrintDocumentPosition(pos wasm.DocumentPosition) {
	switch pos {
	case wasm.DocumentPositionDisconnected:
		fmt.Println("DocumentPositionDisconnected")
	case wasm.DocumentPositionPreceding:
		fmt.Println("DocumentPositionPreceding")
	case wasm.DocumentPositionFollowing:
		fmt.Println("DocumentPositionFollowing")
	case wasm.DocumentPositionContains:
		fmt.Println("DocumentPositionContains")
	case wasm.DocumentPositionContainedBy:
		fmt.Println("DocumentPositionContainedBy")
	case wasm.DocumentPositionImplementationSpecific:
		fmt.Println("DocumentPositionImplementationSpecific")
	}
}

func TestLoginForm(doc wasm.Document) {
	form := wasm.NewHTMLFormElement()
	lblInput := wasm.NewHTMLLabelElement()
	input := wasm.NewHTMLInputElement("text")
	lblPassword := wasm.NewHTMLLabelElement()
	password := wasm.NewHTMLInputElement("password")
	submit := wasm.NewHTMLInputElement("submit")

	form.SetAttribute("action", "#")
	form.SetAttribute("method", "get")
	form.SetAttribute("accept-charset", "utf-8")

	lblInput.SetAttribute("for", "input1")
	lblInput.SetInnerHTML("Input1 Label")
	form.AppendChild(lblInput)

	input.SetAttribute("name", "input1")
	input.SetValue("")
	input.SetId("input1")
	form.AppendChild(input)

	lblPassword.SetAttribute("for", "password")
	lblPassword.SetInnerHTML("Password Label")
	form.AppendChild(lblPassword)

	password.SetAttribute("name", "password")
	password.SetValue("")
	password.SetId("password")
	form.AppendChild(password)

	submit.SetValue("SUBMIT")
	form.AppendChild(submit)

	doc.Body().AppendChild(form)
}
