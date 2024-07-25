"use client";

import { Navbar } from "@/components/main-navbar";
import { Button, buttonVariants } from "@/components/ui/button";
import { FeaturesSectionDemo } from "@/components/ui/features-section";
import { Link } from "lucide-react";
import * as React from "react";

export default function Home() {
  return (
    <main>
      <Navbar className="border-2 border-stone-300 rounded-full"></Navbar>
      <div className="flex justify-center flex-col mt-80 ">
        <div>
          <h1 className="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-7xl text-center">
            My Fitness Aibou
          </h1>
        </div>
        <div>
          <h2 className="scroll-m-20  pb-2 text-4xl font-semibold tracking-tight first:mt-3 text-center">
            All your fitness related needs in one place
          </h2>
        </div>
        <div className="flex justify-center mt-5">
          <Button className="mr-5">Get Started</Button>
          <Button variant="outline">Contact Me</Button>
        </div>
      </div>
      <div className="flex justify-center flex-col mt-56 ">
        <FeaturesSectionDemo></FeaturesSectionDemo>
      </div>
    </main>
  );
}
