"use client";
import React from "react";
import NavBar from "@/components/Navbar";
import { useSession } from "next-auth/react";
import GalleryComponent from "@/components/Gallery";
import SideBarLandingPage from "@/components/SideBarLandingPage";
import SearchBar from "@/components/SearchBar";
import userService from "@/services/user";
import axios from "axios";
<<<<<<< HEAD
import { useEffect, useState } from "react";
import { Gallery } from "@/types";
import photographerGalleryService from "@/services/photographerGalleries";
import Landing from "@/components/Landing";

export default function Home() {

  const [listOfGalleries, setListOfGalleries] = useState<Gallery[]>([]);

  // useEffect(() => {
  //   const fetchAllGalleries = async () => {
  //     const response = await photographerGalleryService.getAllGalleries();
  //     setListOfGalleries(response)
  //   }

  //   fetchAllGalleries();
  // }, [])

  
  const handleOnClick = () => {
    axios.post("/api/auth/signout");
  };
  // This will list of GalleryComponent that get from backend
  // Usage of the `getAllGalleries` function

  return (
    <main>
      <Landing/>

        {/* This is section of displayed gallery view */}
        {/* <div className="mx-5 z-20 w-4/5">
          <div className=" grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
            
            <GalleryComponent
              GalleryName={"White House"}
              Images={["/images/image1.jpg"]}
              Photographer={"Michelle Obama"}
              Price={1290}
            />
            <GalleryComponent
              GalleryName={"American's right"}
              Images={["/images/image1.jpg", "/images/image2.jpg"]}
              Photographer={"Joel Biden"}
              Price={1599}
            />
            <GalleryComponent
              GalleryName={"Golfing In the Forest"}
              Images={[
                "/images/image1.jpg",
                "/images/image2.jpg",
                "/images/image3.jpg",
              ]}
              Photographer={"Donale Trump"}
              Price={1000}
            />
            <GalleryComponent
              GalleryName={"Civil War 2019"}
              Images={["/images/image1.jpg", "/images/image3.jpg"]}
              Photographer={"Osama Bin Laden"}
              Price={1000}
            />
            <GalleryComponent
              GalleryName={"Running Marathon"}
              Images={[
                "/images/image1.jpg",
                "/images/image2.jpg",
                "/images/image3.jpg",
              ]}
              Photographer={"David Goggins"}
              Price={1000}
            />
            <GalleryComponent
              GalleryName={"Mushroom House"}
              Images={[
                "/images/image1.jpg",
                "/images/image2.jpg",
                "/images/image3.jpg",
              ]}
              Photographer={"Blue Smurf"}
              Price={1000}
            />
            <GalleryComponent
              GalleryName={"Face Slapping Oscar"}
              Images={[
                "/images/image1.jpg",
                "/images/image2.jpg",
                "/images/image3.jpg",
              ]}
              Photographer={"Wil Smiths"}
              Price={1000}
            />
            <GalleryComponent
              GalleryName={"The Chinese way"}
              Images={[
                "/images/image1.jpg",
                "/images/image2.jpg",
                "/images/image3.jpg",
              ]}
              Photographer={"Xi Jing Ping"}
              Price={1000}
            />
            <GalleryComponent
              GalleryName={"Suspicious Powder"}
              Images={[
                "/images/image1.jpg",
                "/images/image2.jpg",
                "/images/image3.jpg",
              ]}
              Photographer={"Some Policeman"}
              Price={1000}
            />
            <GalleryComponent
              GalleryName={"Woof Woof Woof"}
              Images={[
                "/images/image1.jpg",
                "/images/image2.jpg",
                "/images/image3.jpg",
              ]}
              Photographer={"Golden Retriever"}
              Price={1000}
            />
            <GalleryComponent
              GalleryName={"Plane"}
              Images={[
                "/images/image1.jpg",
                "/images/image2.jpg",
                "/images/image3.jpg",
              ]}
              Photographer={"Osama Bin Laden"}
              Price={1000}
            />
            <GalleryComponent
              GalleryName={"Handsome Cat"}
              Images={[
                "/images/image1.jpg",
                "/images/image2.jpg",
                "/images/image3.jpg",
              ]}
              Photographer={"Definitely not a cat"}
              Price={1000}
            />
            <GalleryComponent
              GalleryName={"3:33Hrs after midnight"}
              Images={[
                "/images/image1.jpg",
                "/images/image2.jpg",
                "/images/image3.jpg",
              ]}
              Photographer={"Pattapon V."}
              Price={1000}
            />
            <GalleryComponent
              GalleryName={"Giant Water Bottle"}
              Images={[
                "/images/image1.jpg",
                "/images/image2.jpg",
                "/images/image3.jpg",
              ]}
              Photographer={"a guy that sits in front of me"}
              Price={1000}
            />
          </div>
          
        </div> */}
=======
import { useSession } from "next-auth/react";
import { useEffect } from "react";

export default function Home() {
  const { data: session } = useSession();

  const fetchUserInfo = async () => {
    const response = await userService.getMyUserInfo();
    console.log(response);
  };
  fetchUserInfo();
  return (
    <main>
      <NavBar />
      <div className="">Pic-Keeper</div>
>>>>>>> d1017c362207e2fb219c1c18fdb73b3a8b6e11bc
    </main>
  );
}
