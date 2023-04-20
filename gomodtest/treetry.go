package main

import "fmt"

type Menu struct {
	Key      string
	Order    string
	Link     string
	SubMenus []Menu
}

type MenuApi struct {
	ID            uint64
	Key           string
	Order         string
	Link          string
	ParentMenusID uint64
	HeaderID      string
}

var cacheMap = func(ms []*MenuApi) map[uint64][]*MenuApi {
	mmap := make(map[uint64][]*MenuApi)
	for _, v := range ms {
		if v.ParentMenusID != 0 {
			mmap[v.ParentMenusID] = append(mmap[v.ParentMenusID], v)
		}
	}
	return mmap
}(getAllMenu("1"))

func getParentNode(ms []*MenuApi) (sli []*MenuApi) {
	//mmap := make(map[*MenuApi]struct{})
	for _, v := range ms {
		if v.ParentMenusID == 0 {
			sli = append(sli, v)
		}
	}
	return
}

func getChildNode(m *MenuApi) []*MenuApi {
	return cacheMap[m.ID]
}

func getAllMenu(headerID string) []*MenuApi {
	return []*MenuApi{
		// 3 个一级菜单
		{
			ID:            1,
			Key:           "一级菜单 1",
			ParentMenusID: 0,
			HeaderID:      headerID,
		},
		{
			ID:            2,
			Key:           "一级菜单 2",
			ParentMenusID: 0,
			HeaderID:      headerID,
		},
		{
			ID:            3,
			Key:           "一级菜单 3",
			ParentMenusID: 0,
			HeaderID:      headerID,
		},
		// 一级菜单 1 的 3 个子菜单
		{
			ID:            4,
			Key:           "一级菜单 1-1",
			ParentMenusID: 1,
			HeaderID:      headerID,
		},
		{
			ID:            5,
			Key:           "一级菜单 1-2",
			ParentMenusID: 1,
			HeaderID:      headerID,
		},
		{
			ID:            6,
			Key:           "一级菜单 1-3",
			ParentMenusID: 1,
			HeaderID:      headerID,
		},
		// 一级菜单 2 的 3 个子菜单
		{
			ID:            7,
			Key:           "一级菜单 2-1",
			ParentMenusID: 2,
			HeaderID:      headerID,
		},
		{
			ID:            8,
			Key:           "一级菜单 2-2",
			ParentMenusID: 2,
			HeaderID:      headerID,
		},
		{
			ID:            9,
			Key:           "一级菜单 2-3",
			ParentMenusID: 2,
			HeaderID:      headerID,
		},
		// 一级菜单 3 的 3 个子菜单
		{
			ID:            10,
			Key:           "一级菜单 3-1",
			ParentMenusID: 3,
			HeaderID:      headerID,
		},
		{
			ID:            11,
			Key:           "一级菜单 3-2",
			ParentMenusID: 3,
			HeaderID:      headerID,
		},
		{
			ID:            12,
			Key:           "一级菜单 3-3",
			ParentMenusID: 3,
			HeaderID:      headerID,
		},
	}
}

func deep(m *Menu, nodes []*MenuApi) {
	for i, n := range nodes {
		cn := getChildNode(n)
		if len(cn) == 0 {
			m.SubMenus = append(m.SubMenus, *convMenuApiToMenu(n))
			continue
		}
		deep(&m.SubMenus[i], cn)
	}
}

func convMenuApiToMenu(mm *MenuApi) *Menu {
	return &Menu{
		Key:      mm.Key,
		Order:    mm.Order,
		Link:     mm.Link,
		SubMenus: []Menu{},
	}
}

func Range(headerID string) *Menu {
	var allMenu = getAllMenu(headerID)
	var n = getParentNode(allMenu)
	var m = new(Menu)
	for _, v := range n {
		m.SubMenus = append(m.SubMenus, *convMenuApiToMenu(v))
	}
	deep(m, n)
	fmt.Println(n)
	return m
}

func main() {
	for _, menu := range Range("1").SubMenus {
		fmt.Printf("%+v\n", menu)
	}
}

func printPtrSlice(ms []*MenuApi) {
	for _, v := range ms {
		fmt.Printf("%+v\n", *v)
	}
}
