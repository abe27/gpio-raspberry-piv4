/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    env: {
        APP_NAME: process.env.APP_NAME,
        API_HOST: process.env.API_HOST,
        LIMIT_PAGE: process.env.LIMIT_PAGE,
    },
};

module.exports = nextConfig;
