namespace go idlmanagement

service IDLManagement {
    string GetThriftFile(1: string fileName)
}