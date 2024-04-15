use serde::{Deserialize, Serialize};
use sqlx::PgConnection;

#[derive(Serialize, Deserialize)]
pub struct User {
    pub id: i32,
    pub name: String,
    pub email: String,
}

pub struct CreateUserPayload {
    pub name: String,
    pub email: String,
}

pub async fn create_user(
    conn: &mut PgConnection,
    payload: CreateUserPayload,
) -> Result<User, sqlx::Error> {
    sqlx::query_as!(
        User,
        r#"insert into users (name, email) values ($1, $2) returning id, name, email"#,
        payload.name,
        payload.email
    )
    .fetch_one(conn)
    .await
}

pub async fn read_user(conn: &mut PgConnection, id: i32) -> Result<User, sqlx::Error> {
    sqlx::query_as!(
        User,
        r#"select id, name, email from users where id = $1"#,
        id
    )
    .fetch_one(conn)
    .await
}
