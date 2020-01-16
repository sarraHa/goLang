package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    walk(t, ch)
    close(ch)
}

func walk(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }
    walk(t.Left, ch)
    ch <- t.Value
    walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	//panic("not implemented")
	
	// on definit 2 chanel pour stocker les valeaurs des 2 arbres
	ch_t1 := make(chan int) 
	ch_t2 := make(chan int)

	go Walk( t1, ch_t1) 
	go Walk( t2, ch_t2) 

	fmt.Printf("len : %d\n", len( ch_t1) )
	
	
	for {
			v1, ok1 := <-ch_t1
			v2, ok2 := <-ch_t2
		
			if( !ok1 || !ok2){

				return ok1 == ok2
			}
			if( v1 != v2 ){
				return false
			}

			
			
			
	}

	return true	 

}

func main() {
	
	//ch := make(chan int) 
	//go Walk(tree.New(1), ch)
	//for i:= range ch {
	//	fmt.Printf("%d\n", i)
	//}

	if Same(tree.New(1), tree.New(1)) {
		fmt.Printf("OK")
	}else {
		fmt.Printf("FAILED")
	}
	

	
}

	


