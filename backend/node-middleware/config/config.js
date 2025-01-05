const mysql = require('mysql');

// MySQL connection
const connectMySQL = () => {
  const db = mysql.createConnection({
    host: process.env.MYSQL_HOST,
    user: process.env.MYSQL_USER,
    password: process.env.MYSQL_PASSWORD,
    database: process.env.MYSQL_DB
  });

  db.connect((err) => {
    if (err) {
      console.error('MySQL connection error: ', err.stack);
    } else {
      console.log('Connected to MySQL!');
    }
  });
};

module.exports = { connectMySQL };
