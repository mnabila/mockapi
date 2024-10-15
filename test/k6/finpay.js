import http from "k6/http";
import { Counter } from "k6/metrics";

const success = new Counter("success");
const failed = new Counter("failed");

export const options = {
  vus: 100,
  duration: "1m",
};

export default function() {
  const payload = {
    customer: {
      email: "hajar.finnet@gmail.com",
      firstName: "Hajar",
      lastName: "Ismail",
      mobilePhone: "+6281286288844",
    },
    order: {
      id: "1680580078605",
      amount: 200,
      description: "Testing",
      item: [
        {
          name: "name",
          quantity: "1",
          category: "apa",
          description: "apalah",
          unitPrice: "100",
        },
        {
          name: "name",
          quantity: "1",
          category: "apa",
          description: "apalah",
          unitPrice: "100",
        },
      ],
    },
    url: {
      callbackUrl:
        "https://sandbox.finpay.co.id/simdev/finpay/result/resultsuccess.php",
      successUrl:
        "https://sandbox.finpay.co.id/simdev/finpay/result/resultsuccess.php",
      failUrl:
        "https://sandbox.finpay.co.id/simdev/finpay/result/resultfailed.php",
    },
    sourceOfFunds: {
      type: "dana",
    },
  };

  const url = "http://127.0.0.1:8000/pg/payment/card/initiate";

  const resp = http.post(url, JSON.stringify(payload), {
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (resp.status === 200) {
    success.add(1);
  } else {
    failed.add(1);
  }
}
