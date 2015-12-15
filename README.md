#xmldoc

Xmldoc provides a simple xml document model. Such models can be used for contextual processing of xml documents. Xmldoc is not intended to be a replacement for other Golang based xml deserialisation APIs. The model that is generated is 100% mutable so you will need be careful sharing between go routines.

```bash
go get -u github.com/ballpein/xmldoc
```

##Usage

To parse a document supply an io.Reader to xmldoc.Parse( io.Reader ).

```golang

doc,err := xmldoc.Parse(r)
```
NB: The parser will run until it gets an error, so you may get io.EOF back with your document.


To traverse the children directly of any node use the following but be careful for node types such as CDATA where
the Children array will be nil.

```golang
for _,child := range doc.Children {

}
```

For safer traversal use the TraverseChildren function which is available on all nodes and will not fail if the node type does not have children. 
Returning true from the handler function will stop traversal.

```golang
  doc.TraverseChildren(func(node XDNode)(stop bool) {
    ...  
    return stop
    })
```

On element nodes, (XDElement) attributes are available via:

```golang
  for name,value := range ele.Attributes { 
  ...
  }
```

Alternativly on the XDElement there is the attribute traversal function. Returning true from the handler will also stop traversal.

```golang
ele.TraverseAttributes(func(XDName, string) (stop bool) {
...
return stop
})
```

##Getting element value
Depending on how it is interpreted the value of an element is the sum of it's child nodes. In simple cases if you are interested in getting the element value it is likely you are interested in the string value of body of the element.

The following will extract the text value of the body of an XDElement node.

```golang
func textValue(node xmldoc.XDNode)string {
	if ele,ok := node.(*xmldoc.XDElement);  ok {
		b := &bytes.Buffer{}
		ele.TraverseChildren(func(child xmldoc.XDNode)(stop bool) {
			if child.GetType() == xmldoc.CDataType {
				b.Write(child.(*xmldoc.XDCData).Data)
			}
			return stop
		})
		return b.String()
	} else {
		return ""
	}
}
```

##Bugs / Feature requests
Please lodge an issue.
