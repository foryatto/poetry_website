// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package model




// Dynasty is the golang structure for table dynasty.
type Dynasty struct {
    Id   string `orm:"id"   json:"id"`   //   
    Name string `orm:"name" json:"name"` //   
    Time int    `orm:"time" json:"time"` //   
}


// Poem is the golang structure for table poem.
type Poem struct {
    Id      string `orm:"id"      json:"id"`      //   
    Form    string `orm:"form"    json:"form"`    //   
    Tags    string `orm:"tags"    json:"tags"`    //   
    Name    string `orm:"name"    json:"name"`    //   
    Content string `orm:"content" json:"content"` //   
    PoetId  string `orm:"poet_id" json:"poetId"`  //   
}


// Poet is the golang structure for table poet.
type Poet struct {
    Id        string `orm:"id"         json:"id"`        //   
    Name      string `orm:"name"       json:"name"`      //   
    Profile   string `orm:"profile"    json:"profile"`   //   
    DynastyId string `orm:"dynasty_id" json:"dynastyId"` //   
}