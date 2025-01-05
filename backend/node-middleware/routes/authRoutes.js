const express = require('express');
const router = express.Router();
const authMiddleware = require('../middlewares/authMiddleware');

// Dummy routes for login and registration
router.post('/login', (req, res) => {
  // Handle login (authentication)
  res.send('Login route');
});

router.post('/register', (req, res) => {
  // Handle registration
  res.send('Registration route');
});

module.exports = router;
