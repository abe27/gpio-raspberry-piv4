import { useEffect, useState } from "react";
import { AddIcon, CheckIcon, CloseIcon } from "@chakra-ui/icons";
import { reDateTime } from "../hooks/greeter";
import { useToast } from "@chakra-ui/react";

const DeviceComponent = () => {
  const toast = useToast();
  const [data, setData] = useState([]);
  const [deviceName, setDeviceName] = useState(null);
  const [pinNumber, setPinNumber] = useState(0);
  const [alertOn, setAlertOn] = useState(0);

  const fetchData = async () => {
    var requestOptions = {
      method: "GET",
      redirect: "follow",
    };

    const res = await fetch(
      `${process.env.API_HOST}/api/v1/device`,
      requestOptions
    );

    if (res.ok) {
      const r = await res.json();
      setData(r.data);
      // console.dir(r.data);
    }
  };

  const saveData = async () => {
    console.log(alertOn);
    if (deviceName === null || deviceName === "") {
      toast({
        title: "กรุณาระบุชื่ออุปกรณ์ด้วย!",
        status: "error",
        duration: 2500,
        isClosable: true,
        position: "top",
      });
    } else if (pinNumber === null || pinNumber <= 0) {
      toast({
        title: "กรุณาระบุเลขที่ PIN ด้วย!",
        status: "error",
        duration: 2500,
        isClosable: true,
        position: "top",
      });
    } else if (alertOn === null || alertOn <= 0) {
      toast({
        title: "กรุณาระบุอุณหภูมิที่ต้องการแจ้งเตือนด้วย!",
        status: "error",
        duration: 2500,
        isClosable: true,
        position: "top",
      });
    } else {
      var myHeaders = new Headers();
      myHeaders.append("Content-Type", "application/json");

      var raw = JSON.stringify({
        name: deviceName,
        on_pin: parseInt(pinNumber),
        alert_on: parseInt(alertOn),
        is_active: true,
      });

      var requestOptions = {
        method: "POST",
        headers: myHeaders,
        body: raw,
        redirect: "follow",
      };

      const res = await fetch(
        `${process.env.API_HOST}/api/v1/device`,
        requestOptions
      );
      if (res.ok) {
        const r = await res.json();
        toast({
          title: "บันทึกข้อมูลเรียบร้อยแล้ว!",
          status: "success",
          duration: 2500,
          isClosable: true,
          position: "top",
          onCloseComplete: () => fetchData(),
        });
      }

      if (!res.ok) {
        const r = await res.json();
        toast({
          title: r.message,
          status: "error",
          duration: 2500,
          isClosable: true,
          position: "top",
          onCloseComplete: () => fetchData(),
        });
      }
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
            <div className="">
              <div className="form-control w-full">
                <label className="label">
                  <span className="label-text">ชื่ออุปกรณ์</span>
                  <span className="label-text-alt">ระบุข้อมูล</span>
                </label>
                <input
                  type="text"
                  placeholder="Type here"
                  className="input input-sm input-bordered w-full"
                  defaultValue={deviceName}
                  onChange={(e) => setDeviceName(e.target.value)}
                />
                <label className="label">
                  <span className="label-text-alt"></span>
                  <span className="label-text-alt"></span>
                </label>
              </div>
              <div className="flex justify-start space-x-4">
                <div className="form-control w-full">
                  <label className="label">
                    <span className="label-text">เลขที่ PIN</span>
                    <span className="label-text-alt">ระบุข้อมูล</span>
                  </label>
                  <input
                    type="number"
                    placeholder="Type here"
                    className="input input-sm input-bordered w-full"
                    defaultValue={pinNumber}
                    onChange={(e) => setPinNumber(e.target.value)}
                  />
                  <label className="label">
                    <span className="label-text-alt"></span>
                    <span className="label-text-alt"></span>
                  </label>
                </div>
                <div className="form-control w-full">
                  <label className="label">
                    <span className="label-text">แจ้งเตือนที่อุณหภูมิ</span>
                    <span className="label-text-alt">ระบุข้อมูล</span>
                  </label>
                  <input
                    type="number"
                    placeholder="Type here"
                    className="input input-sm input-bordered w-full"
                    defaultValue={alertOn}
                    onChange={(e) => setAlertOn(e.target.value)}
                  />
                  <label className="label">
                    <span className="label-text-alt"></span>
                    <span className="label-text-alt"></span>
                  </label>
                </div>
              </div>
            </div>
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
              <th>อุปกรณ์</th>
              <th>เลขที่ PIN</th>
              <th>แจ้งเตือนที่</th>
              <th>สถานะ</th>
              <th>แกไขล่าสุด</th>
            </tr>
          </thead>
          <tbody>
            {data?.map((i, x) => (
              <tr key={i.id}>
                <th>{x + 1}</th>
                <td>{i.name}</td>
                <td>{i.on_pin}</td>
                <td>
                  <span className="text-yellow-800 hover:cursor-pointer">
                    {i.alert_on}
                  </span>
                </td>
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

export default DeviceComponent;
