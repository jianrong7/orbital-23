namespace go service1v1

struct AddRequest {
    1: required i64 first
    2: required i64 second
}

struct AddResponse {
    1: i64 sum
}

service Service1 {
    AddResponse Add(1: AddRequest req)
}