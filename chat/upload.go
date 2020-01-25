package main

import "net/http"

import "io"

import "io/ioutil"

import "path/filepath"

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid")
	file, header, err := req.FormFile("avatarFile")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	filename := filepath.Join("avatars", userId+filepath.Ext(header.Filename))
	err = ioutil.WriteFile(filename, data, 077)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, "成功")
}
