import MyGallery from "@/components/Gallery/GalleryInfo";

const Home = ({ params }: { params: { galleryId: string } }) => {
  if (!params.galleryId) {
    return <div>No room id specified</div>;
  }
  return (
    <div className="max-w-6xl h-[95vh] m-auto flex items-center justify-center">
      <MyGallery galleryId={params.galleryId} />
    </div>
  );
};

export default Home;
