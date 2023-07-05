#########################
#  Nextjs App S3 Bucket #
#########################
resource "aws_s3_bucket" "frontend_bucket" {
  bucket = "frontend-bucket-${random_uuid.uuid.result}"
}

resource "aws_s3_bucket_website_configuration" "frontend_bucket_website" {
  bucket = aws_s3_bucket.frontend_bucket.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "404.html"
  }
}

resource "aws_s3_bucket_ownership_controls" "frontend_bucket" {
  bucket = aws_s3_bucket.frontend_bucket.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_public_access_block" "frontend_bucket" {
  bucket = aws_s3_bucket.frontend_bucket.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}


resource "aws_s3_bucket_acl" "frontend_bucket" {
  depends_on = [
    aws_s3_bucket_ownership_controls.frontend_bucket,
    aws_s3_bucket_public_access_block.frontend_bucket,
  ]

  bucket = aws_s3_bucket.frontend_bucket.id
  acl    = "public-read"
}

resource "aws_s3_object" "nextjs_app_files" {
  for_each = fileset("${path.module}/../frontend/out", "**/*")

  bucket       = aws_s3_bucket.frontend_bucket.id
  key          = each.value
  source       = "${path.module}/../frontend/out/${each.value}"
  etag         = filemd5("${path.module}/../frontend/out/${each.value}")
  content_type = "text/html"
}

resource "aws_s3_bucket_cors_configuration" "frontend_bucket" {
  bucket = aws_s3_bucket.frontend_bucket.id

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET"]
    allowed_origins = ["*"]
    max_age_seconds = 3000
  }
}

data "aws_iam_policy_document" "public_bucket_access" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.frontend_bucket.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = ["*"]
    }
  }
}

resource "aws_s3_bucket_policy" "this" {
  bucket = aws_s3_bucket.frontend_bucket.id
  policy = data.aws_iam_policy_document.public_bucket_access.json
}
