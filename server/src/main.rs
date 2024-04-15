use server::app::{db::get_pool, routes::routes};
use server::env::{load_env, EnvVar};
use server::utils::logger::setup_logger;
use tokio::net::TcpListener;

#[tokio::main]
async fn main() {
    setup_logger();
    load_env();
    let pool = get_pool().await;
    let app = routes(pool);
    let addr = format!("127.0.0.1:{}", EnvVar::Port.get());
    let listener = TcpListener::bind(addr).await.unwrap();
    tracing::info!("Listening on {}", listener.local_addr().unwrap());
    axum::serve(listener, app).await.unwrap();
}
