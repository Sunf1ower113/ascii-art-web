package server

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	asciiart "ascii-art-web/ascii-art"
	artERrr "ascii-art-web/ascii-art/utils"
)

const (
	OK                  = "OK"
	MethodNotAllowed    = "Method Not Allowed"
	BadRequest          = "Bad Request"
	InternalServerError = "Internal Server Error"
	NotFound            = "Not Found"

	PostFormError = "Post Form Error"
)

type Handler struct {
	mux *http.ServeMux
}

func New(s *http.ServeMux) *Handler {
	h := Handler{s}
	h.registerRoutes()
	return &h
}

func logging(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(
			"%s %s\n",
			r.Method,
			r.URL,
		)
		h.ServeHTTP(w, r)
	})
}

func (h *Handler) registerRoutes() {
	files := http.FileServer(http.Dir("static"))
	h.mux.Handle("/static/", http.StripPrefix("/static/", files))

	h.mux.Handle("/", logging(h.Home))
	h.mux.Handle("/ascii-art", logging(h.AsciiArt))
}

func (h *Handler) redirectError(w http.ResponseWriter, err string, code int) {
	log.Println(err)
	tmpl, err1 := template.New("").ParseFiles("templates/error.html", "templates/base.html")
	if err1 != nil {
		log.Println(err1)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	err = fmt.Sprintf("%d %s", code, err)
	err2 := tmpl.ExecuteTemplate(w, "base", err)
	if err2 != nil {
		log.Println(err2)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.redirectError(w, MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		h.redirectError(w, NotFound, http.StatusNotFound)
		return
	}
	tmpl, err := template.New("").ParseFiles("templates/base.html", "templates/home.html", "templates/asciiart.html")
	if err != nil {
		log.Println(err)
		h.redirectError(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err)
		h.redirectError(w, InternalServerError, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("%s", MethodNotAllowed)
		h.redirectError(w, MethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ascii-art" {
		log.Printf("%s", NotFound)
		h.redirectError(w, NotFound, http.StatusNotFound)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Printf("%s", err.Error())
		h.redirectError(w, BadRequest, http.StatusBadRequest)
		return
	}
	log.Println(r.PostForm)
	for k := range r.PostForm {
		if !(k == "plaintext" || k == "fonts") {
			log.Println(PostFormError)
			h.redirectError(w, BadRequest, http.StatusBadRequest)
			return
		}
	}
	text := r.FormValue("plaintext")
	if len(text) > 400 {
		log.Println(PostFormError)
		h.redirectError(w, BadRequest, http.StatusBadRequest)
		return
	}
	banner := "ascii-art/" + r.FormValue("fonts") + ".txt"
	art, err := asciiart.AsciiArt(text, banner)
	if err != nil {
		log.Printf("%s", err.Error())
		if errors.Is(err, fs.ErrNotExist) {
			h.redirectError(w, InternalServerError, http.StatusInternalServerError)
		} else if errors.Is(err, artERrr.ArtError.InputError) {
			h.redirectError(w, BadRequest, http.StatusBadRequest)
		} else if errors.Is(err, artERrr.ArtError.BannerError1) {
			h.redirectError(w, InternalServerError, http.StatusInternalServerError)
		} else if errors.Is(err, artERrr.ArtError.BannerError2) {
			h.redirectError(w, InternalServerError, http.StatusInternalServerError)
		} else if errors.Is(err, artERrr.ArtError.BannerError3) {
			h.redirectError(w, InternalServerError, http.StatusInternalServerError)
		}
		return
	}
	tmpl, err := template.New("").ParseFiles("templates/home.html", "templates/asciiart.html", "templates/base.html")
	if err != nil {
		h.redirectError(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", art)
	if err != nil {
		h.redirectError(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	log.Printf("%d %s\n", http.StatusOK, OK)
}
