import React, { useEffect, useState } from "react";
import Highcharts from "highcharts";
import {
  HighchartsProvider,
  HighchartsChart,
  Chart,
  XAxis,
  YAxis,
  Title,
  Subtitle,
  Legend,
  LineSeries,
} from "react-jsx-highcharts";

const LineChart = ({ data }) => {
  const [dataA, setDataA] = useState([]);
  const [dataB, setDataB] = useState([]);
  const [plotOptions, setPlotOptions] = useState({ series: {} });
  const [serveA, setServeA] = useState(null);
  const [serveB, setServeB] = useState(null);

  useEffect(() => {
    console.dir(data);
    let a = null;
    let b = null;
    let dbA = [];
    let dbB = [];
    data.map((i) => {
      if (a === null) {
        a = i.device.name;
      } else {
        if (a !== i.device.name) {
          b = i.device.name;
        }
      }

      if (a === i.device.name) {
        dbA.push(i.temp);
        setServeA(i.device.name);
      } else if (b === i.device.name) {
        dbB.push(i.temp);
        setServeB(i.device.name);
      }
      // console.dir(a)
      // console.dir(b)
      setDataA(dbA);
      setDataB(dbB);
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [data]);

  return (
    <div className="app">
      <HighchartsProvider Highcharts={Highcharts}>
        <HighchartsChart plotOptions={plotOptions}>
          <Chart />
          <Title>แสดงข้อมูลอุณหภูมิ</Title>

          {/* <Subtitle></Subtitle> */}

          <Legend layout="vertical" align="right" verticalAlign="middle" />

          <XAxis>
            <XAxis.Title>Time</XAxis.Title>
          </XAxis>

          <YAxis>
            <YAxis.Title>อุณหภูมิ</YAxis.Title>
            <LineSeries name={serveA} data={dataA} />
            <LineSeries name={serveB} data={dataB} />
          </YAxis>
        </HighchartsChart>
      </HighchartsProvider>
    </div>
  );
};

export default LineChart;
