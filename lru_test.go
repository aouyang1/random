package random

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	var val string
	var err error

	size := 4
	lru := NewLRU(size)
	lru.Put("a", "apple")
	if lru.Size() != 1 {
		t.Errorf("expected lru size of 1, but got, %d", lru.Size())
		return
	}
	t.Log(lru.Elements())
	// a

	val, err = lru.Get("a")
	if err = checkValErr(val, err, "apple"); err != nil {
		t.Error(err)
		return
	}
	t.Log(lru.Elements())
	// a

	lru.Put("b", "bear")
	if lru.Size() != 2 {
		t.Errorf("expected lru size of 2, but got, %d", lru.Size())
		return
	}
	t.Log(lru.Elements())
	// b a

	lru.Put("c", "corn")
	if lru.Size() != 3 {
		t.Errorf("expected lru size of 3, but got, %d", lru.Size())
		return
	}
	t.Log(lru.Elements())
	// c b a

	lru.Put("d", "deer")
	if lru.Size() != 4 {
		t.Errorf("expected lru size of 4, but got, %d", lru.Size())
		return
	}
	t.Log(lru.Elements())
	// d c b a

	val, err = lru.Get("a")
	if err = checkValErr(val, err, "apple"); err != nil {
		t.Error(err)
		return
	}
	t.Log(lru.Elements())
	// a d c b

	val, err = lru.Get("b")
	if err = checkValErr(val, err, "bear"); err != nil {
		t.Error(err)
		return
	}
	t.Log(lru.Elements())
	// b a d c

	val, err = lru.Get("c")
	if err = checkValErr(val, err, "corn"); err != nil {
		t.Error(err)
		return
	}
	t.Log(lru.Elements())
	// c b a d

	val, err = lru.Get("d")
	if err = checkValErr(val, err, "deer"); err != nil {
		t.Error(err)
		return
	}
	t.Log(lru.Elements())
	// d c b a

	// exceeding lru size
	lru.Put("e", "elephant")
	if lru.Size() != 4 {
		t.Errorf("expected lru size of 4, but got, %d", lru.Size())
		return
	}
	t.Log(lru.Elements())
	// e d c b

	val, err = lru.Get("a")
	if err = checkValErr(val, err, "apple"); err == nil {
		t.Errorf("expected 'a' to not exist, but does")
		return
	}

	val, err = lru.Get("e")
	if err = checkValErr(val, err, "elephant"); err != nil {
		t.Error(err)
		return
	}
	t.Log(lru.Elements())
	// e d c b

	val, err = lru.Get("b")
	if err = checkValErr(val, err, "bear"); err != nil {
		t.Error(err)
		return
	}
	t.Log(lru.Elements())
	// b e d c

	lru.Put("f", "fairy")
	if lru.Size() != 4 {
		t.Errorf("expected lru size of 4, but got, %d", lru.Size())
		return
	}
	t.Log(lru.Elements())
	// f b e d

	val, err = lru.Get("c")
	if err = checkValErr(val, err, "corn"); err == nil {
		t.Error("expected 'c' to not exist, but does")
		return
	}

	val, err = lru.Get("d")
	if err = checkValErr(val, err, "deer"); err != nil {
		t.Error(err)
		return
	}
	t.Log(lru.Elements())
	// d f b e

	val, err = lru.Get("b")
	if err = checkValErr(val, err, "bear"); err != nil {
		t.Error(err)
		return
	}
	t.Log(lru.Elements())
	// b d f e

}

func checkValErr(val string, err error, exp string) error {
	if err != nil {
		return err
	}
	if val != exp {
		return fmt.Errorf("expected %q but got %q\n", exp, val)
	}
	return nil
}
