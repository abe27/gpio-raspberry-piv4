import { useEffect, useState } from "react";
import { AddIcon, CheckIcon, CloseIcon } from "@chakra-ui/icons";
import { reDateTime } from "../hooks/greeter";

const LineTokenComponent = () => {
  const [data, setData] = useState([]);

  const fetchData = async () => {
    var requestOptions = {
      method: "GET",
      redirect: "follow",
    };

    const res = await fetch(
      `${process.env.API_HOST}/api/v1/token`,
      requestOptions
    );

    if (res.ok) {
      const r = await res.json();
      setData(r.data);
      console.dir(r.data);
    }
  };

  const saveData = () => {
    console.dir("save data");
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
              <th>
                <a
                  href="#addNewDevice"
                  htmlFor="addNewDevice"
                  className="btn btn-ghost btn-xs"
                >
                  <AddIcon color={`blue.500`} />
                </a>
              </th>
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
                <td>{i.description}</td>
                <td>{i.token}</td>
                <td>
                  <button className="btn btn-ghost btn-xs">
                    {i.is_active ? (
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

export default LineTokenComponent;
