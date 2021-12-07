package langserver

import (
	"context"
	"io/ioutil"
	lsp "luahelper-lsp/langserver/protocol"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestProjectDefine(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	paths, _ := filepath.Split(filename)

	strRootPath := paths + "../testdata/define"
	strRootPath, _ = filepath.Abs(strRootPath)

	strRootURI := "file://" + strRootPath
	lspServer := createLspTest(strRootPath, strRootURI)
	context := context.Background()

	fileName := strRootPath + "/" + "test1.lua"
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		t.Fatalf("read file:%s err=%s", fileName, err.Error())
	}

	openParams := lsp.DidOpenTextDocumentParams{
		TextDocument: lsp.TextDocumentItem{
			URI:  lsp.DocumentURI(fileName),
			Text: string(data),
		},
	}
	err1 := lspServer.TextDocumentDidOpen(context, openParams)
	if err1 != nil {
		t.Fatalf("didopen file:%s err=%s", fileName, err1.Error())
	}

	onePosition := lsp.Position{
		Line:      33,
		Character: 4,
	}
	resultRange := lsp.Range{
		Start: lsp.Position{
			Line:      29,
			Character: 6,
		},
		End: lsp.Position{
			Line:      29,
			Character: 9,
		},
	}

	defineParams := lsp.TextDocumentPositionParams{
		TextDocument: lsp.TextDocumentIdentifier{
			URI: lsp.DocumentURI(fileName),
		},
		Position: onePosition,
	}

	resLocationList, err2 := lspServer.TextDocumentDefine(context, defineParams)
	if err2 != nil {
		t.Fatalf("define error")
	}
	if len(resLocationList) != 1 {
		t.Fatalf("location size error")
	}
	res0 := resLocationList[0].Range

	if res0.Start.Line != resultRange.Start.Line || res0.Start.Character != resultRange.Start.Character ||
		res0.End.Line != resultRange.End.Line || res0.End.Character != resultRange.End.Character {
		t.Fatalf("location error")
	}
}

// define 跳转到文件，例如 require("one") 跳转到one/init.lua
func TestProjectDefineFile1(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	paths, _ := filepath.Split(filename)

	strRootPath := paths + "../testdata/define"
	strRootPath, _ = filepath.Abs(strRootPath)

	strRootURI := "file://" + strRootPath
	lspServer := createLspTest(strRootPath, strRootURI)
	context := context.Background()

	fileName := strRootPath + "/" + "test2.lua"
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		t.Fatalf("read file:%s err=%s", fileName, err.Error())
	}

	openParams := lsp.DidOpenTextDocumentParams{
		TextDocument: lsp.TextDocumentItem{
			URI:  lsp.DocumentURI(fileName),
			Text: string(data),
		},
	}
	err1 := lspServer.TextDocumentDidOpen(context, openParams)
	if err1 != nil {
		t.Fatalf("didopen file:%s err=%s", fileName, err1.Error())
	}

	onePosition := lsp.Position{
		Line:      0,
		Character: 27,
	}


	resultRange := lsp.Range{
		Start: lsp.Position{
			Line:      0,
			Character: 0,
		},
		End: lsp.Position{
			Line:      0,
			Character: 1,
		},
	}

	defineParams := lsp.TextDocumentPositionParams{
		TextDocument: lsp.TextDocumentIdentifier{
			URI: lsp.DocumentURI(fileName),
		},
		Position: onePosition,
	}

	resLocationList, err2 := lspServer.TextDocumentDefine(context, defineParams)
	if err2 != nil {
		t.Fatalf("define error")
	}
	if len(resLocationList) != 1 {
		t.Fatalf("location size error")
	}
	strSuffix := "init.lua"
	if !strings.HasSuffix(string(resLocationList[0].URI), strSuffix) {
		t.Fatalf("define file error")
	}

	res0 := resLocationList[0].Range

	if res0.Start.Line != resultRange.Start.Line || res0.Start.Character != resultRange.Start.Character ||
		res0.End.Line != resultRange.End.Line || res0.End.Character != resultRange.End.Character {
		t.Fatalf("location error")
	}
}

// define 跳转到文件，例如 require("one") 跳转到one.lua
func TestProjectDefineFile2(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	paths, _ := filepath.Split(filename)

	strRootPath := paths + "../testdata/define"
	strRootPath, _ = filepath.Abs(strRootPath)

	strRootURI := "file://" + strRootPath
	lspServer := createLspTest(strRootPath, strRootURI)
	context := context.Background()

	fileName := strRootPath + "/" + "test3.lua"
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		t.Fatalf("read file:%s err=%s", fileName, err.Error())
	}

	openParams := lsp.DidOpenTextDocumentParams{
		TextDocument: lsp.TextDocumentItem{
			URI:  lsp.DocumentURI(fileName),
			Text: string(data),
		},
	}
	err1 := lspServer.TextDocumentDidOpen(context, openParams)
	if err1 != nil {
		t.Fatalf("didopen file:%s err=%s", fileName, err1.Error())
	}

	onePosition := lsp.Position{
		Line:      0,
		Character: 27,
	}

	resultRange := lsp.Range{
		Start: lsp.Position{
			Line:      0,
			Character: 0,
		},
		End: lsp.Position{
			Line:      0,
			Character: 1,
		},
	}

	defineParams := lsp.TextDocumentPositionParams{
		TextDocument: lsp.TextDocumentIdentifier{
			URI: lsp.DocumentURI(fileName),
		},
		Position: onePosition,
	}

	resLocationList, err2 := lspServer.TextDocumentDefine(context, defineParams)
	if err2 != nil {
		t.Fatalf("define error")
	}
	if len(resLocationList) != 1 {
		t.Fatalf("location size error")
	}
	strSuffix := "be_define.lua"
	if !strings.HasSuffix(string(resLocationList[0].URI), strSuffix) {
		t.Fatalf("define file error")
	}

	res0 := resLocationList[0].Range

	if res0.Start.Line != resultRange.Start.Line || res0.Start.Character != resultRange.Start.Character ||
		res0.End.Line != resultRange.End.Line || res0.End.Character != resultRange.End.Character {
		t.Fatalf("location error")
	}
}

// define 跳转到文件，例如 import("one.lua") 跳转到one.lua
func TestProjectDefineFile3(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	paths, _ := filepath.Split(filename)

	strRootPath := paths + "../testdata/define"
	strRootPath, _ = filepath.Abs(strRootPath)

	strRootURI := "file://" + strRootPath
	lspServer := createLspTest(strRootPath, strRootURI)
	context := context.Background()

	fileName := strRootPath + "/" + "test4.lua"
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		t.Fatalf("read file:%s err=%s", fileName, err.Error())
	}

	openParams := lsp.DidOpenTextDocumentParams{
		TextDocument: lsp.TextDocumentItem{
			URI:  lsp.DocumentURI(fileName),
			Text: string(data),
		},
	}
	err1 := lspServer.TextDocumentDidOpen(context, openParams)
	if err1 != nil {
		t.Fatalf("didopen file:%s err=%s", fileName, err1.Error())
	}

	onePosition := lsp.Position{
		Line:      0,
		Character: 27,
	}

	resultRange := lsp.Range{
		Start: lsp.Position{
			Line:      0,
			Character: 0,
		},
		End: lsp.Position{
			Line:      0,
			Character: 1,
		},
	}

	defineParams := lsp.TextDocumentPositionParams{
		TextDocument: lsp.TextDocumentIdentifier{
			URI: lsp.DocumentURI(fileName),
		},
		Position: onePosition,
	}

	resLocationList, err2 := lspServer.TextDocumentDefine(context, defineParams)
	if err2 != nil {
		t.Fatalf("define error")
	}
	if len(resLocationList) != 1 {
		t.Fatalf("location size error")
	}
	strSuffix := "be_define.lua"
	if !strings.HasSuffix(string(resLocationList[0].URI), strSuffix) {
		t.Fatalf("define file error")
	}

	res0 := resLocationList[0].Range

	if res0.Start.Line != resultRange.Start.Line || res0.Start.Character != resultRange.Start.Character ||
		res0.End.Line != resultRange.End.Line || res0.End.Character != resultRange.End.Character {
		t.Fatalf("location error")
	}
}
