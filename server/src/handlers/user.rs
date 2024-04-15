use axum::{extract::Path, http::StatusCode, Json};
use serde::{Deserialize, Serialize};

use crate::{
    app::db::DatabaseConnection,
    models::user::{create_user, read_user, CreateUserPayload, User},
};

pub async fn get_user_handler(
    DatabaseConnection(mut conn): DatabaseConnection,
    Path(user_id): Path<i32>,
) -> Result<Json<User>, (StatusCode, String)> {
    read_user(&mut *conn, user_id)
        .await
        .map(Json)
        .map_err(|_| (StatusCode::NOT_FOUND, "User not found".to_string()))
}

#[derive(Serialize, Deserialize)]
pub struct CreateUserRequest {
    name: String,
    email: String,
}

pub async fn create_user_handler(
    DatabaseConnection(mut conn): DatabaseConnection,
    Json(payload): Json<CreateUserRequest>,
) -> Result<Json<User>, (StatusCode, String)> {
    create_user(
        &mut *conn,
        CreateUserPayload {
            name: payload.name,
            email: payload.email,
        },
    )
    .await
    .map(Json)
    .map_err(|_| {
        (
            StatusCode::INTERNAL_SERVER_ERROR,
            "Failed to create user".to_string(),
        )
    })
}
