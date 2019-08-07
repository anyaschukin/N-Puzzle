import "github.com/steakknife/bloomfilter"

const (
  maxElements = 100000
  probCollide = 0.0000001
)

bf, err := bloomfilter.NewOptimal(maxElements, probCollide)
if err != nil {
  panic(err)
}

someValue := ... // must conform to hash.Hash64

bf.Add(someValue)
if bf.Contains(someValue) { // probably true, could be false
  // whatever
}

anotherValue := ... // must also conform to hash.Hash64

if bf.Contains(anotherValue) {
  panic("This should never happen")
}

err := bf.WriteFile("1.bf.gz")  // saves this BF to a file
if err != nil {
  panic(err)
}

bf2, err := bloomfilter.ReadFile("1.bf.gz") // read the BF to another var
if err != nil {
  panic(err)
}