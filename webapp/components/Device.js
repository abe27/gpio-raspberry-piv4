import { AddIcon, CheckIcon, CloseIcon } from "@chakra-ui/icons";
import { useToast } from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { reDateTime } from "../hooks/greeter";

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

  const updateData = async (obj) => {
    console.dir(obj);
    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var raw = JSON.stringify({
      name: obj.name,
      on_pin: obj.on_pin,
      alert_on: obj.alert_on,
      is_active: true,
    });

    var requestOptions = {
      method: "PUT",
      headers: myHeaders,
      body: raw,
      redirect: "follow",
    };

    const res = await fetch(
      `${process.env.API_HOST}/api/v1/device/${obj.id}`,
      requestOptions
    );

    if (res.ok) {
      toast({
        title: "บันทึกข้อมูลเรียบร้อยแล้ว",
        status: "success",
        duration: 1500,
        isClosable: true,
        position: "top-right",
        onCloseComplete: () =>fetchData()
      });
    }

    if (!res.ok) {
      toast({
        title: "เกิดข้อผิดพลาด",
        status: "error",
        duration: 3000,
        isClosable: true,
        position: "top",
        onCloseComplete: () =>fetchData()
      });
    }
  };

  const OnPinChange = (i, e) => {
    let x = 0;
    if (parseInt(e.target.value)) {
      x = parseInt(e.target.value);
    }
    i.on_pin = x;
    updateData(i);
  };

  const alertOnChange = (i, e) => {
    let x = 0;
    if (parseInt(e.target.value)) {
      x = parseInt(e.target.value);
    }
    i.alert_on = x;
    updateData(i);
  };

  const nameChange = (i, e) => {
    i.name = e.target.value;
    updateData(i);
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
          <h3 className="text-lg font-bold">เพิ่มรายการ Device ใหม่</h3>
          <div className="py-4">
            <a
              href="#"
              className="absolute btn btn-sm btn-circle right-2 top-2"
            >
              ✕
            </a>
            <div className="">
              <div className="w-full form-control">
                <label className="label">
                  <span className="label-text">ชื่ออุปกรณ์</span>
                  <span className="label-text-alt">ระบุข้อมูล</span>
                </label>
                <input
                  type="text"
                  placeholder="Type here"
                  className="w-full input input-sm input-bordered"
                  defaultValue={deviceName}
                  onChange={(e) => setDeviceName(e.target.value)}
                />
                <label className="label">
                  <span className="label-text-alt"></span>
                  <span className="label-text-alt"></span>
                </label>
              </div>
              <div className="flex justify-start space-x-4">
                <div className="w-full form-control">
                  <label className="label">
                    <span className="label-text">เลขที่ PIN</span>
                    <span className="label-text-alt">ระบุข้อมูล</span>
                  </label>
                  <input
                    type="number"
                    placeholder="Type here"
                    className="w-full input input-sm input-bordered"
                    defaultValue={pinNumber}
                    onChange={(e) => setPinNumber(e.target.value)}
                  />
                  <label className="label">
                    <span className="label-text-alt"></span>
                    <span className="label-text-alt"></span>
                  </label>
                </div>
                <div className="w-full form-control">
                  <label className="label">
                    <span className="label-text">แจ้งเตือนที่อุณหภูมิ</span>
                    <span className="label-text-alt">ระบุข้อมูล</span>
                  </label>
                  <input
                    type="number"
                    placeholder="Type here"
                    className="w-full input input-sm input-bordered"
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
          </div>
          <div className="modal-action">
            <a href="#" className="btn btn-success" onClick={saveData}>
              บันทึกข้อมูล
            </a>
          </div>
        </div>
      </div>
      <div className="overflow-x-auto">
        <table className="table w-full table-zebra table-compact">
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
                <td>
                <input
                    type="text"
                    className="input input-xs"
                    defaultValue={i.name}
                    onChange={(e) => nameChange(i, e)}
                  />
                </td>
                <td>
                  <input
                    type="number"
                    className="w-24 input input-xs"
                    defaultValue={i.on_pin}
                    onChange={(e) => OnPinChange(i, e)}
                  />
                </td>
                <td>
                  <input
                    type="number"
                    className="w-24 input input-xs"
                    defaultValue={i.alert_on}
                    onChange={(e) => alertOnChange(i, e)}
                  />
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
