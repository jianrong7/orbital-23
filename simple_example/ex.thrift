namespace go api

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

service SimpleExample {
    Response Echo(1: Request req)
    AddResponse Add(1: AddRequest req)
}