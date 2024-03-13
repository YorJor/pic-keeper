"use client";
import { Menu, Transition } from "@headlessui/react";
import { Fragment, useState } from "react";
import { RiArrowDropDownLine } from "react-icons/ri";
import { IoSearch } from "react-icons/io5";

const SideBarLandingPage = () => {
  const classNames = (...classes: string[]) =>
    classes.filter(Boolean).join(" ");

  const [selectedOption, setSelectedOption] = useState("Recommended");
  const [selectedGender, setSelectedGender] = useState("All");

  const handleGenderSelect = (gender:string) => {
    setSelectedGender(gender);
  };

  return (
    <div className="pt-10 flex flex-col px-4 space-y-4">
      <div className="font-semibold">Sort By</div>
      <Menu as="div" className="relative inline-block text-left">
        <div className="flex">
          <Menu.Button className="inline-flex w-full justify-center rounded-md py-2 text-sm font-medium bg-gray-100 focus:outline-none focus-visible:ring-2 focus-visible:ring-white/75">
            {selectedOption}
            <RiArrowDropDownLine className="text-xl" />
          </Menu.Button>
        </div>
        <Transition
          as={Fragment}
          enter="transition ease-out duration-100"
          enterFrom="transform opacity-0 scale-95"
          enterTo="transform opacity-100 scale-100"
          leave="transition ease-in duration-75"
          leaveFrom="transform opacity-100 scale-100"
          leaveTo="transform opacity-0 scale-95"
        >
          <Menu.Items className="absolute w-full mt-2 rounded-md bg-white shadow-lg ring-1 ring-black/5 focus:outline-none">
            <div className="px-1 py-1 ">
              {/* By photographer option */}
              <Menu.Item>
                {({ active }) => (
                  <button
                    onClick={() => setSelectedOption("Recommended")}
                    className={classNames(
                      selectedOption == "Recommended"
                        ? "text-amber-500 underline underline-offset-1 "
                        : "",
                      active ? "bg-gray-100" : "",
                      "block px-4 py-2 text-sm text-gray-700 w-full"
                    )}
                  >
                    Recommended
                  </button>
                )}
              </Menu.Item>
              {/* By gallery name option */}
              <Menu.Item>
                {({ active }) => (
                  <button
                    onClick={() => setSelectedOption("Rating")}
                    className={classNames(
                      selectedOption == "Rating"
                        ? "text-amber-500 underline underline-offset-1"
                        : "",
                      active ? "bg-gray-100" : "",
                      "block px-4 py-2 text-sm text-gray-900 w-full"
                    )}
                  >
                    Rating
                  </button>
                )}
              </Menu.Item>
              <Menu.Item>
                {({ active }) => (
                  <button
                    onClick={() => setSelectedOption("Price")}
                    className={classNames(
                      selectedOption == "Price"
                        ? "text-amber-500 underline underline-offset-1"
                        : "",
                      active ? "bg-gray-100" : "",
                      "block px-4 py-2 text-sm text-gray-900 w-full"
                    )}
                  >
                    Price
                  </button>
                )}
              </Menu.Item>
            </div>
          </Menu.Items>
        </Transition>
      </Menu>
      <div className="font-semibold">Filter By</div>
      <div className="">
        <div className="font-semibold">Gender of Photographer</div>
        <div className="grid grid-cols-3 gap-3 py-2">
            <button 
                className={selectedGender === "Male" ? "outline outline-amber-400 rounded-md p-1 text-amber-400":"rounded-md p-1"} 
                onClick={() => handleGenderSelect("Male")}>
                    Male
            </button>
            <button 
                className={selectedGender === "Female" ? "outline outline-amber-400 rounded-md p-1 text-amber-400":"rounded-md p-1"} 
                onClick={() => handleGenderSelect("Female")}>
                    Female
            </button>
            <button 
                className={selectedGender === "All" ? "outline outline-amber-400 rounded-md p-1 text-amber-400":" rounded-md p-1"} 
                onClick={() => handleGenderSelect("All")}>
                    All
            </button>
        </div>
      </div>
    </div>
  );
};

export default SideBarLandingPage;
