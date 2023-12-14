import dotenv from 'dotenv'

dotenv.config()

const {
	PORT,
	NODE_ENV,
	POSTGRES_HOST,
	POSTGRES_PORT,
	POSTGRES_DB,
	POSTGRES_DB_TEST,
	POSTGRES_USER,
	POSTGRES_PASSWORD
} = process.env

export default {
	port: PORT,
	host: POSTGRES_HOST,
	db_port: POSTGRES_PORT,
	database: POSTGRES_DB,
	user: POSTGRES_USER,
	password: POSTGRES_PASSWORD
}

