use axum::{
    body::{Body, Bytes},
    extract::Request,
    http::StatusCode,
    middleware::Next,
    response::{IntoResponse, Response},
};
use http_body_util::BodyExt;
use tracing;

pub async fn log_req_res(
    req: Request,
    next: Next,
) -> Result<impl IntoResponse, (StatusCode, String)> {
    tracing::info!(
        "::{} {} {}",
        req.method(),
        req.uri().path(),
        req.uri().query().unwrap_or_default()
    );
    let (parts, body) = req.into_parts();
    let bytes = buffer_and_log("Request", body).await?;
    let req = Request::from_parts(parts, Body::from(bytes));

    let res = next.run(req).await;

    let (parts, body) = res.into_parts();
    let bytes = buffer_and_log("Response", body).await?;
    let res = Response::from_parts(parts, Body::from(bytes));

    Ok(res)
}

async fn buffer_and_log<B>(direction: &str, body: B) -> Result<Bytes, (StatusCode, String)>
where
    B: axum::body::HttpBody<Data = Bytes>,
    B::Error: std::fmt::Display,
{
    let bytes = match body.collect().await {
        Ok(collected) => collected.to_bytes(),
        Err(err) => {
            return Err((
                StatusCode::BAD_REQUEST,
                format!("failed to read {direction} body: {err}"),
            ));
        }
    };

    if bytes.is_empty() {
        return Ok(bytes);
    }

    if let Ok(body) = std::str::from_utf8(&bytes) {
        tracing::debug!(
            "{direction} body = {:#?}",
            serde_json::from_str::<serde_json::Value>(body).unwrap()
        );
    }

    Ok(bytes)
}
