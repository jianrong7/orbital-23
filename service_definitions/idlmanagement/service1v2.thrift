namespace go service1v2

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

struct SubRequest {
    1: i64 first
    2: i64 second
}

struct SubResponse {
    1: i64 diff
}

service Service1 {
    AddResponse Add(1: AddRequest req)
    SubResponse Sub(1: SubRequest req)
}