import dotenv from "dotenv";

import nodemailer from "nodemailer";
dotenv.config();
// SMTP
const transporter = nodemailer.createTransport({
  service: "gmail",
  auth: {
    user: process.env.GOOGLE_EMAIL,
    pass: process.env.APP_PASSWORD,
  },
});

export default transporter