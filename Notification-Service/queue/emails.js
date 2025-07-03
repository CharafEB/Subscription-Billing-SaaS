import dotenv from "dotenv";
import SendEmail from "../controller/email.js";
import transporter from "./index.js";
import amqp from "amqplib/callback_api.js";
import UserUpdate from "./response.js";
dotenv.config();

// SendEmails: do send emails
async function SendEmails(UsersData) {
  try {
    if (Array.isArray(UsersData)) {
      for (const user of UsersData) {
        await SendEmail(user, transporter);
        UserUpdate(user);
      }
    } else {
      await SendEmail(UsersData, transporter);
      UserUpdate(UsersData);
    }
    console.log("Emails sent successfully");
  } catch (error) {
    console.error("Error sending emails:", error);
  }
}
// Make connaction with RabitMQ
function RabitMQEmail() {
  amqp.connect(process.env.RabitMQ_Link, function (error0, connection) {
    if (error0) {
      throw error0;
    }
    connection.createChannel(function (error1, channel) {
      if (error1) {
        throw error1;
      }

      const queue = "NotifyUser";

      channel.assertQueue(queue, {
        durable: false,
      });

      channel.consume(
        queue,
        async function (msg) {
          try {
            console.log(msg.content.toString());
            const usersData = JSON.parse(msg.content.toString());
            await SendEmails(usersData);
          } catch (err) {
            console.error("Error parsing message or sending emails:", err);
          }
        },
        {
          noAck: true,
        }
      );
    });
  });
}

export default RabitMQEmail;
