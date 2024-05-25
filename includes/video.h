#include <string>

class VideoServer {
    public:
        VideoServer(const std::string& source);

        std::string get_video(const std::string& name);
        

    private:

    std::string& source;


};
