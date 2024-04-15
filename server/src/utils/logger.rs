use tracing_subscriber::prelude::*;
use tracing_subscriber::{fmt, EnvFilter, Registry};

pub fn setup_logger() {
    let filter = EnvFilter::try_from_env("LOG")
        .or_else(|_| EnvFilter::try_new("debug"))
        .unwrap();

    Registry::default().with(filter).with(fmt::layer()).init();
}
