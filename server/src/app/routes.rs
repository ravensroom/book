use super::{db::DatabaseConnection, errors};
use crate::handlers::user::{create_user_handler, get_user_handler};
use axum::{extract::State, http::StatusCode, routing::get, Router};
use sqlx::postgres::PgPool;

pub fn routes(pool: PgPool) -> Router {
    Router::new()
        .route("/", get(root_get_handler).post(root_post_handler))
        .route("/users", axum::routing::post(create_user_handler))
        .route("/users/:id", get(get_user_handler))
        .with_state(pool)
}

async fn root_get_handler(State(pool): State<PgPool>) -> Result<String, (StatusCode, String)> {
    sqlx::query_scalar("select 'OK'")
        .fetch_one(&pool)
        .await
        .map_err(errors::internal_error)
}

async fn root_post_handler(
    DatabaseConnection(mut conn): DatabaseConnection,
) -> Result<String, (StatusCode, String)> {
    sqlx::query_scalar("select 'OK'")
        .fetch_one(&mut *conn)
        .await
        .map_err(errors::internal_error)
}
