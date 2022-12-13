package chord

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/big"
)

type HashId struct {
	id []byte
}

// !Auxilarry function
func Generate_Hash(inp []byte) HashId {
	h := sha1.New()
	h.Write(inp)
	bs := h.Sum(nil)

	hash := HashId{id: bs}
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
func IsIdBetweenRange_RightEnd_Inclusive(key HashId, min HashId, max HashId) bool {
	key_hash := key.id
	min_hash := min.id
	max_hash := max.id
	firstCond := bytes.Compare(min_hash, key_hash) < 0    // -> True if min < key
	seocondCond := bytes.Compare(key_hash, max_hash) <= 0 // -> True if key <= max

	if firstCond && seocondCond {
		return true
	}
	return false
}

func IsIdBetweenRange_RightEnd_Exclusive(key HashId, min HashId, max HashId) bool {
	key_hash := key.id
	min_hash := min.id
	max_hash := max.id

	firstCond := bytes.Compare(min_hash, key_hash) < 0   // -> True if min < key
	seocondCond := bytes.Compare(key_hash, max_hash) < 0 // -> True if key < max

	if firstCond && seocondCond {
		return true
	}
	return false
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
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("Printing []byte as hex: %x \n", b)
	fmt.Printf("Printing []byte (each byte) as decimal: %v \n", b)
	fmt.Printf("\n")
}

func GetHexBasedStringFromBytes(b []byte) string {
	convertedString := hex.EncodeToString(b)
	return convertedString
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

func Get_X_raised_to_power_Y(x *big.Int, y *big.Int) big.Int {
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

func Find_n_Plus_2_ToPower_k_WholeMod_2_ToPower_m(n *big.Int, k *big.Int, m *big.Int) big.Int {
	x := GetBigIntFromIntegers(2)

	z := Get_X_raised_to_power_Y(&x, k)
	s := AddTwoBigInts(n, &z)

	denom := Get_X_raised_to_power_Y(&x, m)
	out := ModOperationTwoBigInts(&s, &denom)
	return out
}

func GenerateHashIdForFingerIndex(n HashId, indexOfPower int) HashId {

	n_int := GetBigIntFromBytes(n.id)
	k_int := GetBigIntFromIntegers(indexOfPower)
	m_int := GetBigIntFromIntegers(m)

	out_int := Find_n_Plus_2_ToPower_k_WholeMod_2_ToPower_m(&n_int, &k_int, &m_int)
	output := make([]byte, spliceElementsCount, spliceElementsCount)
	final_byte_splices := out_int.FillBytes(output)

	var ret_hash_id HashId
	ret_hash_id.id = final_byte_splices
	return ret_hash_id
}
