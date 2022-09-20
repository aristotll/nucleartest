package plane

type Plane interface {
	Move()
	Image() string
}

type RedPlane struct {
	X, Y    int
	ImgPath string
}

func (r *RedPlane) Move() {
	panic("implement me")
}

func (r *RedPlane) Image() string {
	return r.ImgPath
}

func NewRedPlane() (Plane, error) {
	p := &RedPlane{
		X:       300,
		Y:       500,
		ImgPath: "/Users/zz/GolandProjects/plane-war/img/hero.png",
	}
	//img, err := os.Open("img/hero.png")
	//if err != nil {
	//	return nil, errors.New("open plane image error: " + err.Error())
	//}
	//imgDec, err := png.Decode(img)
	//if err != nil {
	//	return nil, errors.New("dec plane image error: " + err.Error())
	//}
	//p.Img = imgDec
	return p, nil
}
