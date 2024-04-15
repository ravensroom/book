use dotenv::dotenv;
use std::env;
use tracing::log::{info, warn};

pub fn load_env() {
    match dotenv() {
        Ok(_) => {
            info!("Loaded environment variables from .env file");
        }
        Err(e) => {
            panic!("Failed to load environment variables from .env file: {}", e);
        }
    }
}

pub enum EnvVar {
    DatabaseUrl,
    Port,
}

pub struct EnvVarConfig {
    key: String,
    default: String,
}

impl EnvVar {
    fn config(&self) -> EnvVarConfig {
        match self {
            EnvVar::DatabaseUrl => EnvVarConfig {
                key: "DATABASE_URL".to_string(),
                default: "postgres://postgres:password@localhost".to_string(),
            },
            EnvVar::Port => EnvVarConfig {
                key: "PORT".to_string(),
                default: "8080".to_string(),
            },
        }
    }

    pub fn get(&self) -> String {
        let EnvVarConfig { key, default } = self.config();
        match env::var(&key) {
            Ok(value) => {
                info!("Using {} for {}", value, key);
                value
            }
            Err(_) => {
                warn!("{} not found. Using default value: {}", key, default);
                default
            }
        }
    }
}
