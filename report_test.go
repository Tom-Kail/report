package report

import "testing"

func TestNewdoc(t *testing.T) {
	var doc Report
	err := doc.Newdoc("demo.doc")
	if err != nil {
		doc.CloseReport()
		t.Errorf("init doc failed")
	} else {
		doc.CloseReport()
		t.Log("init doc OK")
	}
}

func TestWriteHead(t *testing.T) {
	var doc Report
	doc.Newdoc("demo.doc")
	err := doc.WriteHead()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("WriteHead Succeed")
	}
	doc.Doc.Close()
}

func TestWriteTitle1(t *testing.T) {
	var doc Report
	doc.Newdoc("demo.doc")
	err := doc.WriteTitle1("Hello World")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("TestWriteTitle1 Succeed")
	}
	doc.Doc.Close()
}

func TestWriteTitle2(t *testing.T) {
	var doc Report
	doc.Newdoc("demo.doc")
	err := doc.WriteTitle2("Hello World")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("TestWriteTitle2 Succeed")
	}
	doc.Doc.Close()
}

func TestWriteTitle3(t *testing.T) {
	var doc Report
	doc.Newdoc("demo.doc")
	err := doc.WriteTitle3("Hello World")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("TestWriteTitle3 Succeed")
	}
	doc.Doc.Close()
}
func TestWriteBr(t *testing.T) {
	var doc Report
	doc.Newdoc("demo.doc")
	err := doc.WriteBR()
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("TestWriteBr Succeed")
	}
	doc.Doc.Close()
}

func TestWriteText(t *testing.T) {
	var doc Report
	doc.Newdoc("demo.doc")
	err := doc.WriteText("Hello World")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("TestWriteText succeed")
	}
	doc.Doc.Close()
}
func TestWriteTable(t *testing.T) {
	var doc Report
	doc.Newdoc("demo.doc")
	table := [][][]interface{}{{{"aaa"}, {"bbb"}}, {{"a"}, {"b"}}, {{"xxx"}, {"yyyy"}}}
	err := doc.WriteTable(false, table, [][]interface{}{{"Hello"}, {"World"}})
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("TestWriteTable Succeed")
	}
	doc.Doc.Close()
}
func TestWriteImage(t *testing.T) {
	var doc Report
	doc.Newdoc("demo.doc")
	image1 := &Image{"1.png", "offlineWS-102-risk.png", 140.00, 160.00, 21600, 21600}
	image2 := &Image{"2.png", "offlineWS-102-url.png", 140.00, 160.00, 21600, 21600}
	images := []*Image{image1, image2}
	if err := doc.WriteImage(images, false, ""); err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("TestWriteImage Succeed")
	}
	doc.Doc.Close()
}

func TestWriteEndHead(t *testing.T) {
	var doc Report
	doc.Newdoc("demo.doc")
	err := doc.WriteEndHead(true, true, "Hello World")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log("TestWriteEndHead Succeed")
	}
	doc.Doc.Close()
}
