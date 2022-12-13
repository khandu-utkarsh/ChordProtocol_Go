package chord

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/big"
)

type HashId struct {
	Id []byte
}

// !Auxilarry function
func Generate_Hash(inp []byte) HashId {
	h := sha1.New()
	h.Write(inp)
	bs := h.Sum(nil)

	hash := HashId{Id: bs}
	return hash
}

// !Change these before submitting
func GetCurrentProcessIPAddress() string {
	return "AurojitPanda"
}

func GetCurrentProcessPort() string {
	return ":Sucks"
}

// Implement the comparison function
// Consider right hand boundary
// !This fxn is comparing modulo based
func IsIdBetweenRange_RightEnd_Inclusive(key HashId, min HashId, max HashId) bool {
	keyHash := key.Id
	minHash := min.Id
	maxHash := max.Id

	bytesCompOut := bytes.Compare(minHash, maxHash)
	if bytesCompOut == 0 {
		return true
	} else if bytesCompOut == -1 { // min < max
		fc := bytes.Compare(minHash, keyHash) < 0
		sc := bytes.Compare(keyHash, maxHash) <= 0
		return fc && sc
	} else { //min > max
		//!Inter-changed - Swap
		t := maxHash
		maxHash = minHash
		minHash = t

		//!Complement set conditions
		fc := bytes.Compare(minHash, keyHash) <= 0
		sc := bytes.Compare(keyHash, maxHash) < 0
		return !(fc && sc) //!Returning complement of the reuslt
	}
}

// !Comparison is modulo based here
func IsIdBetweenRangeRightEndExclusive(key HashId, min HashId, max HashId) bool {
	keyHash := key.Id
	minHash := min.Id
	maxHash := max.Id

	bytesCompOut := bytes.Compare(minHash, maxHash)
	if bytesCompOut == 0 {
		return !(bytes.Compare(minHash, keyHash) == 0) //!Returning complement of the reuslt
	} else if bytesCompOut == -1 { // min < max
		fc := bytes.Compare(minHash, keyHash) < 0
		sc := bytes.Compare(keyHash, maxHash) < 0
		return fc && sc
	} else { //min > max
		//!Inter-changed - Swap
		t := maxHash
		maxHash = minHash
		minHash = t

		//!Complement set conditions
		fc := bytes.Compare(minHash, keyHash) <= 0
		sc := bytes.Compare(keyHash, maxHash) <= 0
		return !(fc && sc) //!Returning complement of the reuslt
	}
}

// %Printing functions
// Printing functions for big int
func PrintBigInt(n big.Int) {
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Println("Printing big.Int as decimal: ", n.Text(10))
	fmt.Println("Printing big.Int as hex: ", n.Text(16))
	fmt.Println("Printing big.Int as binary: ", n.Text(2))
	fmt.Printf("\n")

}

func PrintBytesSplices(b []byte) {
	fmt.Println()
	fmt.Printf("Printing []byte as hex: %x \n", b)
	fmt.Printf("Printing []byte (each byte) as decimal: %v \n", b)
}

func GetHexBasedStringFromBytes(b []byte) string {
	convertedString := hex.EncodeToString(b)
	return convertedString
}

func GetByteArrayFromString(s string) []byte {
	// TODO : handle error later
	byteArray, _ := hex.DecodeString(s)

	return byteArray
}

func GetBigIntFromBytes(b []byte) big.Int {
	var n_big_int big.Int
	n_big_int.SetBytes(b)
	return n_big_int
}

func GetBigIntFromIntegers(num int) big.Int {
	var n_big_int big.Int
	n_big_int.SetUint64(uint64(num))
	return n_big_int
}

func GetXRaisedToPowerY(x *big.Int, y *big.Int) big.Int {
	var power_out big.Int
	power_out.Exp(x, y, nil)
	return power_out
}

func AddTwoBigInts(x *big.Int, y *big.Int) big.Int {
	var sum big.Int
	sum.Add(x, y)
	return sum
}

// x%y
func ModOperationTwoBigInts(x *big.Int, y *big.Int) big.Int {
	var mod big.Int
	mod.Mod(x, y)
	return mod
}

func FindNPlus2ToPowerKWholeMod2ToPowerM(n *big.Int, k *big.Int, m *big.Int) big.Int {
	x := GetBigIntFromIntegers(2)

	z := GetXRaisedToPowerY(&x, k)
	s := AddTwoBigInts(n, &z)

	denom := GetXRaisedToPowerY(&x, m)
	out := ModOperationTwoBigInts(&s, &denom)
	return out
}

func GenerateHashIdForFingerIndex(n HashId, indexOfPower int) HashId {

	nInt := GetBigIntFromBytes(n.Id)
	kInt := GetBigIntFromIntegers(indexOfPower)
	mInt := GetBigIntFromIntegers(m)

	outInt := FindNPlus2ToPowerKWholeMod2ToPowerM(&nInt, &kInt, &mInt)
	output := make([]byte, spliceElementsCount, spliceElementsCount)
	finalByteSplices := outInt.FillBytes(output)

	var retHashId HashId
	retHashId.Id = finalByteSplices
	return retHashId
}

func IsSameNode(node1 Node, node2 Node) bool {
	return GetHexBasedStringFromBytes(node1.node_id.Id) == GetHexBasedStringFromBytes(node2.node_id.Id)
}
