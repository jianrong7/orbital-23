namespace go service2v1

struct MulRequest {
    1: required i64 first
    2: required i64 second
}

struct MulResponse {
    1: i64 sum
}

struct DivRequest {
    1: required i64 first
    2: required i64 second
}

struct DivResponse {
    1: double ratio
}

service Service2 {
    MulResponse Mul(1: MulRequest req)
    DivResponse Div(1: DivRequest req)
}