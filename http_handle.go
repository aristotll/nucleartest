package main

import (
	"context"
	"io"
	"log"
	"net/http"
)

type HandleImpl struct {
	ctx context.Context
}

func NewHandleImpl(ctx context.Context) *HandleImpl {
	return &HandleImpl{
		ctx: ctx,
	}
}

func (h *HandleImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r = r.WithContext(h.ctx)
	ctxVal := r.Context().Value("k").(string)
	io.WriteString(w, ctxVal)
}

func main() {
	ctx := context.WithValue(context.TODO(), "k", "v")
	h := NewHandleImpl(ctx)
	http.Handle("/", h)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
