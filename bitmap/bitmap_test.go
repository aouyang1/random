package main

import (
	"reflect"
	"testing"
)

func TestContainer(t *testing.T) {
	c := NewContainer()

	c.Insert(3)
	if !reflect.DeepEqual(c.GetAll(), []int{3}) {
		t.Errorf("first insert, %+v", c)
	}
	if len(c.Bitmap) != 1 {
		t.Errorf("expected 1, but got , %d", len(c.Bitmap))
	}

	c.Insert(7)
	if !reflect.DeepEqual(c.GetAll(), []int{3, 7}) {
		t.Errorf("insert first container, %+v", c.GetAll())
	}
	if len(c.Bitmap) != 1 {
		t.Errorf("expected 1, but got , %d", len(c.Bitmap))
	}

	c.Insert(64)
	if !reflect.DeepEqual(c.GetAll(), []int{3, 7, 64}) {
		t.Errorf("append second container, %+v", c.GetAll())
	}
	if len(c.Bitmap) != 2 {
		t.Errorf("expected 2, but got , %d", len(c.Bitmap))
	}

	c.Insert(69)
	if !reflect.DeepEqual(c.GetAll(), []int{3, 7, 64, 69}) {
		t.Errorf("insert second container, %+v", c.GetAll())
	}
	if len(c.Bitmap) != 2 {
		t.Errorf("expected 2, but got , %d", len(c.Bitmap))
	}

	c.Insert(0)
	if !reflect.DeepEqual(c.GetAll(), []int{0, 3, 7, 64, 69}) {
		t.Errorf("insert 0, %+v", c.GetAll())
	}
	if len(c.Bitmap) != 2 {
		t.Errorf("expected 2, but got , %d", len(c.Bitmap))
	}
}

func TestOr(t *testing.T) {
	a := NewContainer()
	a.Insert(1)
	a.Insert(5)
	a.Insert(68)
	a.Insert(129)

	b := NewContainer()
	b.Insert(0)
	b.Insert(7)
	b.Insert(5)
	b.Insert(10048)

	c := a.Copy()
	c.Or(b)
	if !reflect.DeepEqual(c.GetAll(), []int{0, 1, 5, 7, 68, 129, 10048}) {
		t.Errorf("%+v", c.GetAll())
	}
	if len(c.Bitmap) != 158 {
		t.Errorf("expected 158, but got , %d", len(c.Bitmap))
	}

	c = b.Copy()
	c.Or(a)
	if !reflect.DeepEqual(c.GetAll(), []int{0, 1, 5, 7, 68, 129, 10048}) {
		t.Errorf("%+v", c.GetAll())
	}
	if len(c.Bitmap) != 158 {
		t.Errorf("expected 158, but got , %d", len(c.Bitmap))
	}

}

func TestAnd(t *testing.T) {
	a := NewContainer()
	a.Insert(1)
	a.Insert(5)
	a.Insert(68)
	a.Insert(129)

	b := NewContainer()
	b.Insert(0)
	b.Insert(7)
	b.Insert(5)
	b.Insert(10048)

	c := a.Copy()
	c.And(b)
	if !reflect.DeepEqual(c.GetAll(), []int{5}) {
		t.Errorf("%+v", c.GetAll())
	}
	if len(c.Bitmap) != 3 {
		t.Errorf("expected 1, but got , %d", len(c.Bitmap))
	}

	c = b.Copy()
	c.And(a)
	if !reflect.DeepEqual(c.GetAll(), []int{5}) {
		t.Errorf("%+v", c.GetAll())
	}
	if len(c.Bitmap) != 3 {
		t.Errorf("expected 1, but got , %d", len(c.Bitmap))
	}
}
