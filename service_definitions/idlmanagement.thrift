namespace go idlmanagement

service IDLManagement {
    string GetThriftFile(1: string serviceName, 2: string serviceVersion)
}