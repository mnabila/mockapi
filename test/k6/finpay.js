import http from "k6/http";
import { check, fail, sleep } from "k6";

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

  const isSuccess = check(resp, {
    "status code must be 200": (response) => response.status === 200,
  });

  if (!isSuccess) {
    fail(`request failed with status ${resp.status_text}`);
  }
}
