package _9999

import (
	"fmt"
	"github.com/indrariksa/cobapakcage/model"
	"github.com/indrariksa/cobapakcage/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Rumah"
	phonenumber := "68122221814"
	checkin := "masuk"
	biodata := model.Karyawan{
		Nama:         "Drake",
		Phone_number: "628456456211",
		Jabatan:      "Rakyat Biasa",
	}
	hasil := module.InsertPresensi(long, lat, lokasi, phonenumber, checkin, biodata)
	fmt.Println(hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "68122221814"
	biodata, err := module.GetKaryawanFromPhoneNumber(phonenumber, module.MongoConn, "presensi")
	if err != nil {
		t.Fatalf("error calling GetKaryawanFromPhoneNumber: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetPresensiFromID(t *testing.T) {
	id := "641898e76dd7bd217d69762a"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetPresensiFromID(objectID, module.MongoConn, "presensi")
	if err != nil {
		t.Fatalf("error calling GetPresensiFromID: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetAll(t *testing.T) {
	data := module.GetAllPresensi(module.MongoConn, "presensi")
	fmt.Println(data)
}
