import { useEffect, useState } from "react";
import { AddIcon, CheckIcon, CloseIcon } from "@chakra-ui/icons";
import { reDateTime } from "../hooks/greeter";

const NotificationComponent = () => {
  const [data, setData] = useState([]);

  const fetchData = async () => {
    var requestOptions = {
      method: "GET",
      redirect: "follow",
    };

    const res = await fetch(
      `${process.env.API_HOST}/api/v1/notification`,
      requestOptions
    );

    if (res.ok) {
      const r = await res.json();
      setData(r.data);
      // console.dir(r.data);
    }
  };

  const saveData = () => {
    console.dir("save data");
  };

  const updateStatus = async (obj) => {
    obj.is_accept = !obj.is_accept;
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var raw = JSON.stringify({
      device_id: obj.device.name,
      line_token_id: obj.line_token.description,
      is_accept: obj.is_accept,
      is_active: obj.is_active,
    });

    var requestOptions = {
      method: "PUT",
      headers: myHeaders,
      body: raw,
      redirect: "follow",
    };

    const res = await fetch(
      `${process.env.API_HOST}/api/v1/notification/${obj.id}`,
      requestOptions
    );
    if (res.ok) {
      const r = await res.json();
      fetchData();
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <>
      <div className="modal" id="addNewDevice">
        <div className="modal-box">
          <h3 className="font-bold text-lg">เพิ่มรายการ Device ใหม่</h3>
          <p className="py-4">
            <a
              href="#"
              className="btn btn-sm btn-circle absolute right-2 top-2"
            >
              ✕
            </a>
            You ve been selected for a chance to get one year of subscription to
            use Wikipedia for free!
          </p>
          <div className="modal-action">
            <a href="#" className="btn btn-success" onClick={saveData}>
              บันทึกข้อมูล
            </a>
          </div>
        </div>
      </div>
      <div className="overflow-x-auto">
        <table className="table table-zebra table-compact w-full">
          <thead>
            <tr>
              <th></th>
              <th>หัวข้อ</th>
              <th>เลขที่ Token</th>
              <th>สถานะ</th>
              <th>แกไขล่าสุด</th>
            </tr>
          </thead>
          <tbody>
            {data?.map((i, x) => (
              <tr key={i.id}>
                <th>{x + 1}</th>
                <td>{i.device.name}</td>
                <td>{i.line_token.description}</td>
                <td>
                  <button
                    className="btn btn-ghost btn-xs"
                    onClick={() => updateStatus(i)}
                  >
                    {i.is_accept ? (
                      <CheckIcon color="green.500" />
                    ) : (
                      <CloseIcon color={`red.500`} />
                    )}
                  </button>
                </td>
                <td>{reDateTime(i.updated_at)}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </>
  );
};

export default NotificationComponent;
