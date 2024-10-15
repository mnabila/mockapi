import { sleep } from "k6";
import http from "k6/http";

export const options = {
  vus: 10,
  duration: "30s",
};

export default function() {
  const idtrx = Date.now();
  const url = `http://127.0.0.1:8000/api/h2h?id=MA0001&pin=12345&user=MAU001&pass=12345&kodeproduk=SB5&tujuan=085123456789&counter=1&idtrx=${idtrx}&jenis=1`;
  http.get(url.toString());
  sleep(1);
}
