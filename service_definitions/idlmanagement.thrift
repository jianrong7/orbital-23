namespace go idlmanagement

service IDLManagement {
    string CheckVersion()
    string GetServiceThriftFileName(1: string serviceName)
    string GetThriftFile(1: string serviceName)
}