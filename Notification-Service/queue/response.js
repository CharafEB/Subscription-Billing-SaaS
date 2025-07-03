import dotenv from "dotenv";
import amqp from "amqplib/callback_api.js";
dotenv.config();

// Make connaction with RabitMQ
function UserUpdate(msg) {
  amqp.connect(process.env.RabitMQ_Link, function (error0, connection) {
    if (error0) {
      throw error0;
    }
    connection.createChannel(function (error1, channel) {
      if (error1) {
        throw error1;
      }
      var queue = "UpdateUser";

      channel.assertQueue(queue, {
        durable: false,
      });


      channel.sendToQueue(queue, Buffer.from(JSON.stringify(msg)));
      console.log(" [x] Sent %s", JSON.stringify(msg));
    });
  });
}

export default UserUpdate;
