namespace go service2v1

struct Request {
    1: string message
}

struct Response {
    1: string message
}

struct MulRequest {
    1: i64 first
    2: i64 second
}

struct MulResponse {
    1: i64 sum
}

service Service2 {
    MulResponse Mul(1: MulRequest req)
}