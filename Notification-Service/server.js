import express from "express";
import dotenv from "dotenv";
import RabitMQEmail from "./queue/emails.js";
dotenv.config();

RabitMQEmail();
const app = express();
const PORT = process.env.PORT || 3000;

//Run the server
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
