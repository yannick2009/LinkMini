import { toast } from "@steveyuowo/vue-hot-toast";
import axios from "axios";

const serverUrl: string = import.meta.env.VITE_SERVER_URL;

type Url = {
    longUrl: string;
    urlHash: string;
    clickNum: number;
    qrCode: string;
    createdAt: Date;
    ExpiredAt: Date;
}

const requestMethods = {
    post: "POST",
    get: "GET",
};


export const createShortUrl = async (url: string) => {
  try {
    const response = await axios<Url>({
        method: requestMethods.post,
        url: `${serverUrl}/create`,
        data: { url },
    });
    toast.success("Short URL created successfully");
    return response;
  } catch (error) {
    console.error(error);
    toast.error("Failed to create short URL");
  }
}

export const getShortUrlStats = async (urlHash: string) => {
    try {
        const response = await axios<Url>({
            method: requestMethods.get,
            url: `${serverUrl}/stats/${urlHash}`,
        });
        toast.success("Short URL stats fetched successfully");
        return response;
    } catch (error) {
        console.error(error);
        toast.error("Failed to get short URL stats");
    }
    }
