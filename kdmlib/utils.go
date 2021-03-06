package kdmlib

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	K     = 20
	ALPHA = 3
)

//Generates a Random ID, of specified length, given by constant IDLENGTH
//The returned ID is a bitwise representation
func GenerateRandID(seed int64, binLength int) string {
	id := ""
	rand.Seed(time.Now().UnixNano() - seed)
	for i := 0; i < binLength; i++ {
		id += strconv.Itoa(rand.Intn(2))
	}

	return id
}

//Converts a bitwise-represented ID into a HEX-represented ID
func ConvertToHexAddr(binAddr string) string {
	hexAddr := ""
	for i := 0; i < len(binAddr)/4; i++ {
		newPart := []rune(binAddr[(i * 4) : (i*4)+4])
		newIntPart := 0
		for j := 0; j < 4; j++ {
			if string(newPart[j]) == "1" {
				newIntPart += int(math.Pow(2, float64(3-j)))
			}
		}
		hexAddr += strconv.FormatInt(int64(newIntPart), 16)
	}
	return hexAddr
}

//Generates, or converts an ID from HEX-represented input ID
func GenerateIDFromHex(hexAddr string) string {
	binAddr := ""
	hexXAddr := []rune(hexAddr[:])

	for i := 0; i < len(hexXAddr); i++ {
		switch string(hexXAddr[i]) {
		case "0":
			binAddr += "0000"
			break
		case "1":
			binAddr += "0001"
			break
		case "2":
			binAddr += "0010"
			break
		case "3":
			binAddr += "0011"
			break
		case "4":
			binAddr += "0100"
			break
		case "5":
			binAddr += "0101"
			break
		case "6":
			binAddr += "0110"
			break
		case "7":
			binAddr += "0111"
			break
		case "8":
			binAddr += "1000"
			break
		case "9":
			binAddr += "1001"
			break
		case "a":
			binAddr += "1010"
			break
		case "b":
			binAddr += "1011"
			break
		case "c":
			binAddr += "1100"
			break
		case "d":
			binAddr += "1101"
			break
		case "e":
			binAddr += "1110"
			break
		case "f":
			binAddr += "1111"
			break
		}
	}

	return binAddr
}

func GenerateZeroID(binLength int) string {
	zeroString := ""
	for i := 0; i < binLength; i++ {
		zeroString += "0"
	}
	return zeroString
}

//Encodes a file name into a 160 bit ID
//Maximum 19 characters is allowed in fileName.
/*
func HashKademliaID(fileName string) string {
	f := hex.EncodeToString([]byte(fileName))
	if len(f) > 38 {
		fmt.Println("Name of file can be maximum 19 characters, including file extension.")
	}
	f = f + "03"
	for len(f) < 40 {
		f = f + "01"
	}
	return f
}
*/

func ComputeDistance(id1 string, id2 string) (string, error) {
	if len(id1) != len(id2) {
		return "", errors.New("lengths of the IDs are different")
	} else {
		var sb strings.Builder
		for i := 0; i < len(id1); i++ {
			if id1[i] == id2[i] {
				sb.WriteString("0")
			} else {
				sb.WriteString("1")
			}
		}
		return sb.String(), nil
	}
}

func ConvertToUDPAddr(contact AddressTriple) *net.UDPAddr {
	addr, err := net.ResolveUDPAddr("udp", contact.Ip+":"+contact.Port)

	if err != nil {
		log.Fatal("Error: ", err)
		return nil
	} else {
		return addr
	}
}

func AlreadyAsked(asked []AddressTriple, c AddressTriple) bool {
	for _, element := range asked {
		if c.Id == element.Id {
			return true
		}
	}
	return false
}

func PrintListOfContacts(description string, contactList []AddressTriple) {
	fmt.Println(description)
	for index := range contactList {
		fmt.Println("['"+ConvertToHexAddr(contactList[index].Id)+"'", contactList[index].Ip+":"+contactList[index].Port+"]")
	}
}

func IsResultClosed(ch <-chan interface{}) bool {
	select {

	case <-ch:
		return true
	default:
	}

	return false
}

func IsLookupClosed(ch <-chan LookupOrder) bool {
	select {

	case <-ch:
		return true
	default:
	}

	return false
}

/*

func SendPostRequest(triple AddressTriple, content []byte) string {
	req, err := http.NewRequest("POST", "http://"+triple.Ip+":8000/?fromNetwork=true", bytes.NewBuffer(content))
	if err != nil {
		fmt.Println("error creating the post request")
	} else {
		client := &http.Client{}
		client.Timeout = time.Second * 5
		resp, err := client.Do(req)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("error sending the post request")
			return ""
		} else {
			b, _ := ioutil.ReadAll(resp.Body)
			return string(b)
		}
	}
	return ""
}

func SendGetRequest(triple AddressTriple, name string) []byte {
	req, err := http.NewRequest("GET", "http://"+triple.Ip+":8000/"+name, nil)
	if err != nil {
		fmt.Println("error creating the post request")
	} else {
		client := &http.Client{}
		client.Timeout = time.Second * 5
		resp, err := client.Do(req)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("error sending the get request")
			return nil
		} else {
			b, _ := ioutil.ReadAll(resp.Body)
			return b
		}
	}
	return nil
}

*/
