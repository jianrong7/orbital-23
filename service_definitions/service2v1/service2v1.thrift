namespace go service2v1

struct MulRequest {
    1: required i64 first
    2: required i64 second
}

struct MulResponse {
    1: i64 sum
}

service Service2 {
    MulResponse Mul(1: MulRequest req)
}