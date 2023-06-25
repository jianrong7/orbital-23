namespace go service1v1

struct Request {
    1: string message
}

struct Response {
    1: string message
}

struct AddRequest {
    1: i64 first
    2: i64 second
}

struct AddResponse {
    1: i64 sum
}

service Service1 {
    AddResponse Add(1: AddRequest req)
}